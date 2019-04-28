## Install

# Install C dependencies
apt update
apt install -y build-essential pkg-config libpth-dev libavcodec-dev libavutil-dev libavformat-dev libswscale-dev libwebp-dev libopencv-dev libgeoip-dev git
apt-get dist-upgrade -y

# Install Node.js
wget -qO- https://deb.nodesource.com/setup_9.x | bash -
apt-get install -y nodejs

# Install and init PostgreSQL
echo deb http://apt.postgresql.org/pub/repos/apt/ $(lsb_release -sc)-pgdg main >> /etc/apt/sources.list.d/pgdg.list
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
apt update
apt install -y postgresql
service postgresql start
echo "Now run this manually:
su postgres
createuser -P meguca
createdb -T template0 -E UTF8 -O meguca meguca
exit
"

# Increase PostgreSQL connection limit by changing `max_connections` to 1024
# nano /etc/postgresql/11/main/postgresql.conf

# Install Go
wget -O- https://dl.google.com/go/go1.11.9.linux-amd64.tar.gz | tar xpz -C /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
source /etc/profile

# Clone and build meguca
git clone https://github.com/bakape/meguca.git meguca
cd meguca
make

# Run meguca
./meguca help
./meguca -a :80
