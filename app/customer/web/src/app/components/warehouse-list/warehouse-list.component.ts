import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { faWarehouse } from '@fortawesome/free-solid-svg-icons';
import { UtilsService } from 'src/app/services/utils.service';
import { Warehouse } from '../warehouse-card/warehouse';
import * as cloneDeep from 'lodash/cloneDeep';
import { MessageService, MessagePayload, MessageType } from 'src/app/services/message.service';

@Component({
  selector: 'app-warehouse-list',
  templateUrl: './warehouse-list.component.html',
  styleUrls: ['./warehouse-list.component.scss']
})
export class WarehouseListComponent implements OnInit {

  emptyIcon = faWarehouse;
  warehouses: Warehouse[];
  masterList: Warehouse[];
  hasFilter = false;

  constructor(
    private http: HttpClient,
    private utils: UtilsService,
    private msgService: MessageService,
  ) { 
    this.warehouses = []
    this.masterList = []
  }

  ngOnInit(): void {
    this.init();
    this.subscribeEvents();
  }

  async init() {

    let randomNumber = (min:number, max:number) => { 
      return Math.ceil(Math.random() * (max - min) + min);
    } 

    let data = await this.initAssets();
    this.masterList = data.map(wh=> {
      wh.generalInfo.rating = randomNumber(1, 5);
      return wh;
    });

    this.warehouses = cloneDeep(this.masterList);
  }

  async initAssets() : Promise<Warehouse[]> {
    const url = `${this.utils.baseUrl()}/api/warehouse`;
    const res = await this.http.get<any>(url).toPromise();
    return JSON.parse(res.message);
  }

  subscribeEvents() {

    this.msgService.getMessage()
    .subscribe((m: MessagePayload)=>{

      this.hasFilter = true;

      switch (m.type){
        case MessageType.ClearFilter:
          this.onClearFilter()
          break;

        case MessageType.RatingsFilter:
          this.onFilterByRatings(m)
          break;
        
        case MessageType.RateFilter:
          this.onFilterByRateRange(m)
          break;
            
        case MessageType.SpaceFilter:
          this.onFilterBySpace(m);
          break;
        case MessageType.LocationFilter:
          this.onFilterByAddress(m);
          break;
      }
    },
    err =>{
      console.error('msg service error', err);
    })
  }

  onClearFilter() {
    this.hasFilter = false;
    this.warehouses = cloneDeep(this.masterList);
  }

  onFilterByRatings(m: MessagePayload) {
    const result = this.masterList.filter(e=>{
      return e.generalInfo.rating >= m.data;
    });

    this.warehouses = cloneDeep(result);
  }

  onFilterByRateRange(m: MessagePayload) {
    const result = this.masterList.filter(e=>{
      return e.generalInfo.rate >= m.data.from && e.generalInfo.rate <= m.data.to;
    });

    this.warehouses = cloneDeep(result);
  }

  onFilterBySpace(m: MessagePayload) {
    const result = this.masterList.filter(e=>{
      const available = e.generalInfo.totalArea - e.generalInfo.allocatedArea;
      return available >= m.data;
    });

    this.warehouses = cloneDeep(result);
  }

  onFilterByAddress(m: MessagePayload) {
    const word = m.data.toLowerCase();
    const result = this.masterList.filter(e=>{
      const addr = e.generalInfo.address.toLowerCase();
      return addr.includes(word);
    });

    this.warehouses = cloneDeep(result);
  }
}
