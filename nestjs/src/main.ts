import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

// Muitos comentários porque NestJS é uma tecnologia nova para mim
// NestJS trabalha com Arquitetura de para Contêiners de Serviço e Arquitetura Modular
// Tudo que fazemos em NestJS é um módulo
// Esse arquivo main.js inicia o projeto
// O AppModule nesse caso é o módulo raíz da aplicação
async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await app.listen(3000);
}
bootstrap();
