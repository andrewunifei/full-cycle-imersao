import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { PrismaModule } from './prisma/prisma.module';
import { AssetsModule } from './assets/assets.module';
import { WalletsModule } from './wallets/wallets.module';

// Muitos comentários porque NestJS é uma tecnologia nova para mim
// ES7 Decorator - @Module
// Esse Decorator torna AppModule um módulo
// Dentro do módulo podemos registrar artefatos (como controller) que são conhecidos no módulo
// O que é conhecido o NestJS faz automações
// Por exemplo, uma vez que o controller é registrado (AppController),
// o NestJS reconhece e guia a rota através de Decorators,
// então uma classe para virar um controller se coloca o decorator @Controller
// e todas as rotas que você quer habilitar são os verbos HTTP
// que tbm são declarados com Decorators (@Get por exemplo)
@Module({
  imports: [PrismaModule, AssetsModule, WalletsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
