package db

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Mongo struct de conexao para o driver oficial go.mongodb.org/mongo-driver.
// Suporta MongoDB 5.0+ e ReplicaSet de forma nativa.
type Mongo struct {
	client   *mongo.Client
	database *mongo.Database
	config   Config
}

// NewMongo abre uma conexão com o MongoDB usando o driver oficial.
// Se config implementar ReplicaSetConfig,
// os parâmetros de ReplicaSet são aplicados à URI (replicaSet, authSource) e o
// readPreference é configurado programaticamente nas options.
func NewMongo(ctx context.Context, config Config) (*Mongo, error) {
	uri := buildMongoURI(config)

	clientOpts := options.Client().ApplyURI(uri)

	if rs, ok := config.(ReplicaSetConfig); ok {
		if pref := mapReadPreference(rs.GetReadPreference()); pref != nil {
			clientOpts.SetReadPreference(pref)
		}
	}

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, err
	}

	return &Mongo{
		client:   client,
		database: client.Database(config.GetDatabase()),
		config:   config,
	}, nil
}

// TestMongoConnection abre uma conexão, pinga o servidor e desconecta.
// Útil para health checks. 
// O timeout é controlado pelo contexto.
func TestMongoConnection(ctx context.Context, config Config) error {
	m, err := NewMongo(ctx, config)
	if err != nil {
		return err
	}
	defer m.Disconnect(context.Background())
	return nil
}

// Client retorna o *mongo.Client 
// (transactions, sessions, etc.).
func (m *Mongo) Client() *mongo.Client {
	return m.client
}

// Database retorna o *mongo.Database conectado.
func (m *Mongo) Database() *mongo.Database {
	return m.database
}

// Collection é um atalho para m.Database().Collection(name).
func (m *Mongo) Collection(name string) *mongo.Collection {
	return m.database.Collection(name)
}

// Ping verifica se o servidor está acessível.
func (m *Mongo) Ping(ctx context.Context) error {
	if m == nil || m.client == nil {
		return errors.New("mongo client is not initialized")
	}
	return m.client.Ping(ctx, nil)
}

// Disconnect encerra a conexão com o MongoDB.
// Pode ser chamado múltiplas vezes sem efeito colateral. 
// Use um contexto com timeout para limitar o tempo de espera durante o shutdown.
func (m *Mongo) Disconnect(ctx context.Context) error {
	if m == nil || m.client == nil {
		return nil
	}
	return m.client.Disconnect(ctx)
}

// buildMongoURI monta a URI de conexão fazendo URL-encoding de usuário e senha.
// Caracteres especiais como @, #, $ na senha são codificados corretamente —
// requisito do parser de URIs do MongoDB.
func buildMongoURI(config Config) string {
	var userInfo string
	if config.GetUser() != "" || config.GetPassword() != "" {
		userInfo = url.QueryEscape(config.GetUser()) + ":" + url.QueryEscape(config.GetPassword()) + "@"
	}

	uri := "mongodb://" + userInfo + config.GetHost() + ":" + strconv.Itoa(config.GetPort()) + "/" + config.GetDatabase()

	if rs, ok := config.(ReplicaSetConfig); ok {
		var params []string
		if v := rs.GetReplicaSet(); v != "" {
			params = append(params, "replicaSet="+v)
		}
		if v := rs.GetAuthSource(); v != "" {
			params = append(params, "authSource="+v)
		}
		if len(params) > 0 {
			uri += "?" + strings.Join(params, "&")
		}
	}

	return uri
}

// mapReadPreference converte a string da config no tipo do driver oficial.
// Retorna nil se a string for vazia ou desconhecida — o driver usará o default (primary).
func mapReadPreference(pref string) *readpref.ReadPref {
	switch pref {
	case "primary":
		return readpref.Primary()
	case "primaryPreferred":
		return readpref.PrimaryPreferred()
	case "secondary":
		return readpref.Secondary()
	case "secondaryPreferred":
		return readpref.SecondaryPreferred()
	case "nearest":
		return readpref.Nearest()
	default:
		return nil
	}
}

// defaultConnectTimeout é uma constante de referência para callers que queiram
// um timeout razoável de conexão. Não é aplicada automaticamente — quem chama
// NewMongo deve passar o contexto com o timeout desejado.
const defaultConnectTimeout = 10 * time.Second

// DefaultConnectTimeout retorna o timeout padrão sugerido para conexões.
func DefaultConnectTimeout() time.Duration {
	return defaultConnectTimeout
}