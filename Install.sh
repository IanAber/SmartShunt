if [ "$USER" != "root" ]
then
    echo "Please run this as root or with sudo"
    exit 2
fi
systemctl stop SmartShunt
cp dist/amd64/SmartShunt /usr/bin

chmod +x /usr/bin/SmartShunt

if ! test -d "/SmartShunt"; then
  mkdir /SmartShunt
fi
systemctl start SmartShunt

echo "Done"






