set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE SCHEMA auth;
    CREATE SCHEMA users;
    CREATE SCHEMA warehouse;
    CREATE SCHEMA checkout;
    
    -- Opcional: Permisos específicos por esquema
    GRANT ALL PRIVILEGES ON SCHEMA auth TO microservice_user;
    GRANT ALL PRIVILEGES ON SCHEMA users TO microservice_user;
    GRANT ALL PRIVILEGES ON SCHEMA warehouse TO microservice_user;
    GRANT ALL PRIVILEGES ON SCHEMA checkout TO microservice_user;
EOSQL