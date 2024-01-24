if !(type psqldef > /dev/null 2>&1); then
    echo "psqldef が存在しません。 https://github.com/sqldef/sqldef?tab=readme-ov-file#installation"
    exit 1
fi

source .env

case $1 in
    production)
        ENV=production;;
    *)
        ENV=local;;
esac

echo ENV=${ENV}

FILE=./init/1.schema.sql

case $ENV in
    local)
        psqldef -W postgres -f $FILE postgres
        ;;
    production)
        psqldef -h $PRODUCTION_HOST -p $PRODUCTION_PORT -U $PRODUCTION_USER -W $PRODUCTION_PASSWORD -f $FILE $PRODUCTION_DATABASE
        ;;
esac
