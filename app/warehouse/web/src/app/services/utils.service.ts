import { Injectable } from '@angular/core';
import { isDevMode } from '@angular/core';
import {v4 as uuidv4} from 'uuid';

@Injectable({
  providedIn: 'root'
})
export class UtilsService {

  constructor() { }

  baseUrl(): string{
    return isDevMode() ? 'http://localhost:8081' : 'http://localhost:8081';
  }

  getUUID():string {
    return uuidv4();
  }
}
