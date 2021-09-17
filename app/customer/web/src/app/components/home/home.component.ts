import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
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
    this.identity = (await this.initProfile());
  }

  async initProfile() {
    const url = `${this.utils.baseUrl()}/api/identity`;
    const res = await this.http.get<Payload>(url).toPromise();
    return JSON.parse(res.message);
  }

  onOpenWallet() {

  }
  
  onShowTransactions() {

  }
}


interface Identity {
  Name: string;
}
 
interface Payload {
  message: string;
}