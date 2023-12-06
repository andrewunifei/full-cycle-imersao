import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';

// Exemplo de classe que é um controller
// Declarada com o Decorator @Controller
// O AppService é a classe que o NestJS recomenda para as regras de negócio, acesso ao banco de dados,
// que serão servidas no caminho fornecido nos verbos HTTP
@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }
}
