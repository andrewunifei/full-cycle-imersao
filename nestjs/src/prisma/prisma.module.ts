import { Global, Module } from '@nestjs/common';
import { PrismaService } from './prisma.service';

// Módulo para que o Prisma esteja disponível no NestJS
@Global()
@Module({
  providers: [PrismaService],
  exports: [PrismaService],
})
export class PrismaModule {}
