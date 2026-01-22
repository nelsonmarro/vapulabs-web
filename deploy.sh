#!/bin/bash

# Configuración
USER="nelson"
IP="82.25.95.39"
APP_DIR="/opt/naphsoft"
STATIC_DIR="/var/www/naphsoft"

echo "--- 🚀 Iniciando despliegue de NaphSoft ---"

# 1. Compilar todo localmente
echo "📦 Compilando aplicación..."
make build
if [ $? -ne 0 ]; then
    echo "❌ Error en la compilación. Abortando."
    exit 1
fi

# 2. Detener servicio y limpiar binario antiguo para evitar "text file busy"
echo "🛑 Deteniendo servicio para actualizar binario..."
ssh -t $USER@$IP "sudo systemctl stop naphsoft.service && sudo rm -f $APP_DIR/server"

# 3. Subir el binario y asegurar permisos
echo "📤 Subiendo nuevo binario a $APP_DIR..."
scp bin/server $USER@$IP:$APP_DIR/server
ssh $USER@$IP "chmod +x $APP_DIR/server"

# 4. Sincronizar carpeta static (solo sube cambios, ignorando descargas pesadas)
echo "📂 Sincronizando archivos estáticos en $STATIC_DIR..."
rsync -avz --delete --exclude 'downloads/' static/ $USER@$IP:$STATIC_DIR/static/

# 5. Subir configuración .env
echo "⚙️  Subiendo configuración .env..."
scp .env $USER@$IP:$APP_DIR/.env

# 6. Reiniciar el servicio
echo "🔄 Reiniciando servicio naphsoft..."
ssh -t $USER@$IP "sudo systemctl reset-failed naphsoft.service && sudo systemctl start naphsoft.service"

echo "--- ✅ ¡Despliegue completado con éxito! ---"