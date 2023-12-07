import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class WalletAssetsService {
    constructor(private prismaService: PrismaService) {}

    all(filter: { wallet_id: string }) {
        return this.prismaService.walletAsset.findMany({
            where: {
                wallet_id: filter.wallet_id,
            },
            include: { // Disponibiliza acesso aos registros relacionados
                Asset: {
                    select: {
                        id: true,
                        ticker: true,
                        price: true,
                    },
                },
            },
        })
    }

    create(data: { wallet_id: string; asset_id: string; shares: number }) {
        return this.prismaService.walletAsset.create({
            data,
        })
    }
}

