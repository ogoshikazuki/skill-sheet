if [ $# != 1 ]; then
    echo バージョンを指定してください。（例：v1.0.0）
    exit 1
fi

SQLDEF_VERSION=$1

curl -OL https://github.com/k0kubun/sqldef/releases/download/${SQLDEF_VERSION}/psqldef_linux_amd64.tar.gz
sudo tar xf psqldef_linux_amd64.tar.gz -C /usr/local/bin/
rm psqldef_linux_amd64.tar.gz
