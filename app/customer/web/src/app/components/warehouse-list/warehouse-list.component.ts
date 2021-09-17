import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { faWarehouse } from '@fortawesome/free-solid-svg-icons';
import { UtilsService } from 'src/app/services/utils.service';
import { Warehouse } from '../warehouse-card/warehouse';

@Component({
  selector: 'app-warehouse-list',
  templateUrl: './warehouse-list.component.html',
  styleUrls: ['./warehouse-list.component.scss']
})
export class WarehouseListComponent implements OnInit {

  emptyIcon = faWarehouse;

  warehouses: Warehouse[];

  constructor(
    private http: HttpClient,
    private utils: UtilsService,
  ) { 
    this.warehouses = []
  }

  ngOnInit(): void {

    this.init();
  }

  async init() {

    let randomNumber = (min:number, max:number) => { 
      return Math.ceil(Math.random() * (max - min) + min);
    } 

    let data = await this.initAssets();
    this.warehouses = data.map(wh=> {
      wh.generalInfo.rating = randomNumber(1, 5);
      return wh;
    });
  }

  async initAssets() : Promise<Warehouse[]> {
    const url = `${this.utils.baseUrl()}/api/warehouse`;
    const res = await this.http.get<any>(url).toPromise();
    return JSON.parse(res.message);
  }
}
