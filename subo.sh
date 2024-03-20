# Comandos Mac
git add . 
git commit -m "Next commit"
git push

# Construir el binario ejecutable
go build -tags lambda.norpc -o bootstrap main.go

# Empaquetar el binario en un archivo ZIP
zip main.zip bootstrap

# Limpiar archivos temporales (el binario)
rm bootstrap