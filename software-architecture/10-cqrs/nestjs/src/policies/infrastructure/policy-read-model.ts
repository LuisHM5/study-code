import { Injectable } from '@nestjs/common';

// Vista de lectura: plana, con la forma que necesita la pantalla.
export interface PolicyView {
  id: string;
  number: string;
  startDate: string;
  endDate: string;
}

// Lado de LECTURA: un almacén aparte del modelo de escritura.
@Injectable()
export class PolicyReadModel {
  private readonly views = new Map<string, PolicyView>();

  upsert(view: PolicyView): void {
    this.views.set(view.id, view);
  }

  findById(id: string): PolicyView | undefined {
    return this.views.get(id);
  }

  expiringBetween(from: string, to: string): PolicyView[] {
    return [...this.views.values()]
      .filter((v) => v.endDate >= from && v.endDate <= to)
      .sort((a, b) => a.endDate.localeCompare(b.endDate));
  }
}
