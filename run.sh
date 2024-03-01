# ## Start pgadmin4
gunicorn \
    --daemon \
    --bind 0.0.0.0:5050 \
    --workers=1 \
    --threads=4 \
    --chdir /Users/hoangle/.local/share/virtualenvs/server_go-EghkzA9f/lib/python3.9/site-packages/pgadmin4 \
    pgAdmin4:app
