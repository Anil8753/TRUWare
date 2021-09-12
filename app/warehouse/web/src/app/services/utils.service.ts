import { Injectable } from '@angular/core';
import { isDevMode } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class UtilsService {

  constructor() { }

  baseUrl(): string{
    return isDevMode() ? 'http://localhost:8081' : 'http://localhost:8081';
  }
}
