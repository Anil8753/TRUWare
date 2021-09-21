import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { isDevMode } from '@angular/core';
import {v4 as uuidv4} from 'uuid';

@Injectable({
  providedIn: 'root'
})
export class UtilsService {

  constructor(
    private http: HttpClient,
  ) { }

  baseUrl(): string{
    return isDevMode() ? 'http://localhost:8082' : 'http://localhost:8082';
  }

  getUUID():string {
    return uuidv4();
  }

  async getWallet(): Promise<WalletEntry> {

    try {
      const url = `${this.baseUrl()}/api/wallet`;
      const res = await this.http.get<any>(url).toPromise();
      return JSON.parse(res.message) as WalletEntry;
    } catch (e){
      console.error(e);
    }

    return null
  }

  async getRegistration(): Promise<RegistrationData> {

    try {
      const url = `${this.baseUrl()}/api/registeration`;
      const res = await this.http.get<any>(url).toPromise();
      return JSON.parse(res.message) as RegistrationData;
    } catch (e){
      console.error(e);
    }
  
    return null
  }
}

export interface WalletEntry {
  owner: string;
  balance: number;
  refNo: string;
}

export class RegistrationData {

  id: string;
  name: string;
  address: string;
  contact: string;
  email: string;
  gst:string;
}
 