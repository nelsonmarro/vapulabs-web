# Guía de Despliegue en VPS (Arch Linux) - NaphSoft

Esta guía documenta el proceso completo para configurar un VPS con Arch Linux desde cero, desplegar la aplicación web NaphSoft (Go + Templ), configurar Nginx como proxy inverso, asegurar el sitio con SSL y desplegar la documentación de Docusaurus.

## 1. Configuración Inicial del VPS

### Acceso y Creación de Usuario

1. **Acceso SSH como root:**

   ```bash
   ssh root@<TU_IP_VPS>
   ```

2. **Actualizar el sistema:**

   ```bash
   pacman -Sy archlinux-keyring --noconfirm
   pacman -key --init
   pacman-key --populate archlinux
   pacman -Syu --noconfirm
   ```

3. **Instalar herramientas base:**

   ```bash
   pacman -S sudo git base-devel nano rsync ufw --noconfirm
   ```

4. **Crear usuario no-root:**

   ```bash
   useradd -m -G wheel nelson
   passwd nelson
   ```

5. **Configurar Sudoers:**

   ```bash
   EDITOR=nano visudo
   # Descomentar la línea: %wheel ALL=(ALL:ALL) ALL
   # Añadir al final para reinicios sin contraseña:
   # nelson ALL=(ALL) NOPASSWD: /usr/bin/systemctl restart naphsoft
   ```

### Firewall (UFW)

```bash
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow ssh
sudo ufw allow http
sudo ufw allow https
sudo ufw enable
```

---

## 2. Instalación del Stack Web

### Nginx y Certbot

```bash
sudo pacman -S nginx certbot certbot-nginx --noconfirm
```

### Estructura de Directorios

Organizamos el servidor en dos ubicaciones principales:

- `/opt/naphsoft`: Binarios y configuración de la aplicación (Backend).
- `/var/www/naphsoft`: Archivos estáticos, descargas y documentación (Frontend/Static).

```bash
sudo mkdir -p /opt/naphsoft
sudo mkdir -p /var/www/naphsoft/static/downloads
sudo mkdir -p /var/www/naphsoft/docs
sudo chown -R nelson:nelson /opt/naphsoft
sudo chown -R nelson:nelson /var/www/naphsoft
```

---

## 3. Configuración de Nginx

### `nginx.conf` Principal

Editar `/etc/nginx/nginx.conf`:

```nginx
worker_processes  1;
events {
    worker_connections  1024;
}
http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  65;
    types_hash_max_size 4096;

    # Importar configuraciones de sitios
    include sites-enabled/*;
}
```

### Configuración del Sitio (`/etc/nginx/sites-available/naphsoft`)

```nginx
# Web Principal (Go App)
server {
    listen 80;
    server_name naphsoft.dev www.naphsoft.dev;

    # Proxy a la App de Go
    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Optimización para descargas pesadas (Instalador)
    location /static/downloads/ {
        alias /var/www/naphsoft/static/downloads/;
        sendfile on;
        tcp_nopush on;
        tcp_nodelay on;
    }
}

# Documentación (Docusaurus)
server {
    listen 80;
    server_name docs.verith.naphsoft.dev;

    root /var/www/naphsoft/docs;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

**Activar el sitio:**

```bash
sudo mkdir -p /etc/nginx/sites-enabled
sudo ln -s /etc/nginx/sites-available/naphsoft /etc/nginx/sites-enabled/
sudo systemctl enable --now nginx
```

---

## 4. Despliegue de la Aplicación (Go + Templ)

### Servicio Systemd

Crear `/etc/systemd/system/naphsoft.service`:

```ini
[Unit]
Description=NaphSoft Web App
After=network.target

[Service]
Type=simple
User=nelson
WorkingDirectory=/opt/naphsoft
ExecStart=/opt/naphsoft/server
Restart=always
EnvironmentFile=/opt/naphsoft/.env
Environment=STATIC_DIR=/var/www/naphsoft/static

[Install]
WantedBy=multi-user.target
```

Activar: `sudo systemctl enable --now naphsoft.service`

### Script de Despliegue (`deploy.sh`)

Script local para automatizar actualizaciones:

```bash
#!/bin/bash
# ... (Compilación local)
# Subida de binario a /opt/naphsoft
# Subida de assets a /var/www/naphsoft/static (excluyendo downloads)
# Reinicio del servicio
```

---

## 5. Subida del Instalador (Archivo Grande)

Para subir `Verith_Setup.zip` (300MB+) sin que `rsync` lo borre en cada despliegue:

1. Asegurar que `deploy.sh` tenga `--exclude 'downloads/'` en el comando rsync.
2. Subir manualmente una vez:

   ```bash
   scp /ruta/local/Verith_Setup.zip nelson@<IP>:/var/www/naphsoft/static/downloads/
   ```

---

## 6. Despliegue de Documentación (Docusaurus)

1. Generar build local: `npm run build`
2. Sincronizar con el VPS:

   ```bash
   rsync -avz --delete build/ nelson@<IP>:/var/www/naphsoft/docs/
   ```

---

## 7. Seguridad SSL (HTTPS)

Una vez que los DNS (Registros A) apunten a la IP del VPS:

```bash
sudo certbot --nginx -d naphsoft.dev -d www.naphsoft.dev -d docs.verith.naphsoft.dev
sudo systemctl enable --now certbot-renew.timer
```

## Resumen de Comandos Útiles

- **Desplegar App Web:** `./deploy.sh` (desde local)
- **Ver Logs App:** `ssh nelson@<IP> "journalctl -u naphsoft -f"`
- **Reiniciar Nginx:** `ssh nelson@<IP> "sudo systemctl restart nginx"`
- **Subir Docs:** `rsync -avz --delete build/ nelson@<IP>:/var/www/naphsoft/docs/`
