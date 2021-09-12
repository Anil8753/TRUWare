import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { AssetComponent } from '../asset/asset.component';
import { Asset, GeneralInfo, Postion } from 'src/app/components/asset/asset';
import { HttpClient } from '@angular/common/http';
import { UtilsService } from 'src/app/services/utils.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  asset: Asset;
  identity: Identity;

  constructor(
    private http: HttpClient,
    private modalService: NgbModal,
    private utils: UtilsService,
  ) { }

  ngOnInit(): void {
    this.init();
  }

  async init() {
    try {
      this.identity = await this.initProfile();
      this.asset = await this.initAssets();
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
    const url = `${this.utils.baseUrl()}/api/warehouse/52`;
    const res = await this.http.get<Payload>(url).toPromise();
    return JSON.parse(res.message);
  }

  onClick() {

    const modalRef = this.modalService.open(AssetComponent, { backdrop: 'static'});
    modalRef.componentInstance.name = 'World';

    modalRef.result
    .then((asset: Asset)=>{
      if (asset)
        this.asset = asset;
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