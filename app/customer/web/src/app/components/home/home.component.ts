import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { HttpClient } from '@angular/common/http';
import { UtilsService } from 'src/app/services/utils.service';
import { faWallet, faMoneyCheck, faUser } from '@fortawesome/free-solid-svg-icons';
import * as cloneDeep from 'lodash/cloneDeep';
import { WalletComponent } from '../wallet/wallet.component';
import { TransactionsComponent } from '../transactions/transactions.component';
import { AccountComponent } from '../account/account.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  walletIcon = faWallet;
  txnIcon = faMoneyCheck;
  profileIcon = faUser;

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
    this.modalService.open(WalletComponent, { backdrop:'static'});
  }
  
  onShowTransactions() {
    this.modalService.open(TransactionsComponent, { backdrop:'static', size: 'lg'});
  }

  onUserProfile() {
    this.modalService.open(AccountComponent, { backdrop:'static'});
  }
}


interface Identity {
  Name: string;
}
 
interface Payload {
  message: string;
}