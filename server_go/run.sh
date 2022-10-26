## Start pgadmin4
gunicorn  --bind 0.0.0.0:5050 \
          --workers=1 \
          --threads=4 \
          --chdir /home/hoang/.local/share/virtualenvs/server_go-d0vzcw5v/lib/python3.9/site-packages/pgadmin4 \
          pgAdmin4:app