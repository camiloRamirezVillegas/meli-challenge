#!/bin/bash

# Function to generate random string
generate_secret_key() {
    < /dev/urandom tr -dc A-Za-z0-9 | head -c${1:-32}
}

# Create .env file
cat <<EOF > .env
# Environment variables
SECRET_KEY=$(generate_secret_key)
DATABASE_URL=your_database_url_here
DEBUG=True
EOF

echo ".env file created successfully with a randomly generated SECRET_KEY."