
export class GeneralInfo {
    name: string;
    phone: string;
    email: string;
    address: string;
    details: string;
    rate: number;
    totalArea: number;
    allocatedArea: number;
  }
  
  export class Postion {
    latitude: number;
    longitude: number;
  }
  
  export class Asset {
    id: string;
    status: number;
    generalInfo: GeneralInfo;
    allocations: any[];
    postion: Postion;
  }