import { Injectable } from '@nestjs/common';

// @Injectable é o Decorator que significa que a classe (serviço nesse caso)
// pode ser injetada em Controllers e em outros serviços também
// Mas é necessário registrar o serviço no Módulo dentro da propriedade "providers"
@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello NestJS, I\'m in love!';
  }
}
