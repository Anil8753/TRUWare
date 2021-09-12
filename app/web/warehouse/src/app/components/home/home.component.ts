import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { AssetComponent } from '../asset/asset.component';
import { Asset, GeneralInfo, Postion } from 'src/app/components/asset/asset';
import { HttpClient } from '@angular/common/http';
import { UtilsService } from 'src/app/services/utils.service';
import { faEdit, faWarehouse } from '@fortawesome/free-solid-svg-icons';
import * as cloneDeep from 'lodash/cloneDeep';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  faEdit = faEdit;
  iconWarehouseAdd = faWarehouse;

  selected: Asset;
  assets: Asset[];
  identity: Identity;

  constructor(
    private http: HttpClient,
    private modalService: NgbModal,
    private utils: UtilsService,
  ) { }

  ngOnInit(): void {
    this.assets = [];
    this.init();
  }

  async init() {
    try {
      this.identity = await this.initProfile();
      this.assets = await this.initAssets();

      if (this.assets.length > 0) {
        this.selected = this.assets[0];
      }

    } catch (e) {
      console.error(e);
    }
  }

  async initProfile() {
    const url = `${this.utils.baseUrl()}/api/warehouse/identity`;
    const res = await this.http.get<Payload>(url).toPromise();
    return JSON.parse(res.message);
  }

  async initAssets() {
    const url = `${this.utils.baseUrl()}/api/warehouse`;
    const res = await this.http.get<Payload>(url).toPromise();
    return JSON.parse(res.message);
  }

  onChange(index: number) {
    this.selected = this.assets[index];
  }

  onAddNew() {

    const modalRef = this.modalService.open(AssetComponent, { backdrop: 'static'});
    modalRef.componentInstance.edit = false;

    modalRef.result
    .then((asset: Asset)=>{
      if (!!asset) {
        if (this.assets.length === 0) {
          this.selected = asset;
        }
        this.assets.push(asset);
      }
    })
    .catch(e=>{
      alert(JSON.stringify(e));
    });
  }

  onUpdate() {
    
    const modalRef = this.modalService.open(AssetComponent, { backdrop: 'static'});
    modalRef.componentInstance.asset = cloneDeep(this.selected);;
    modalRef.componentInstance.edit = true;

    modalRef.result
    .then((asset: Asset)=>{
      if (!!asset) {
        const index = this.assets.findIndex(ws=> ws.id == asset.id);
        if (index >= 0 && index < this.assets.length)
        {
          this.assets[index] = asset;
          this.selected = asset;
          console.log(asset);
        }
      }
    })
    .catch(e=>{
      alert(JSON.stringify(e));
    });
  }
}


interface Identity {
  Name: string;
}
 
interface Payload {
  message: string;
}