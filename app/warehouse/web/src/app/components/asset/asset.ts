
export class GeneralInfo {
    gst: string;
    name: string;
    phone: string;
    email: string;
    address: string;
    details: string;
    rate: number;
    panalityPremature: number;
    panalityAfterLimit: number;
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