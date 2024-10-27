# Usa la imagen oficial de Go 1.21.6
FROM golang:1.21.6

# Establece el directorio de trabajo en /app
WORKDIR /app

# descarga  los modulos
COPY go.mod go.sum ./
RUN go mod download

COPY . .
EXPOSE 8080
# Compila la aplicación
RUN go build -o newsletter
ENV \
    mongoUrl="mongodb+srv://gustavohdzmz:welcome55@newsletter.9soh00l.mongodb.net/?retryWrites=true&w=majority&appName=newsletter" \
    mongoDb="newsletter-app" \
    mongoNewsletterCollection="newsletters" \
    mongoSubscriberCollection="subscribers" \
    emailSender="postmaster@sandbox2671e155d0ea48bfaff044f1e55edfb4.mailgun.org" \
    emailPass="e25f082ec2536527b01fc10f114989ee-784975b6-92a16f66" \
    smtpServer="smtp.mailgun.org" \
    smtpPort=587

# Ejecuta la aplicación al iniciar el contenedor
CMD ["./newsletter"]