import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import { UtilsService } from 'src/app/services/utils.service';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss']
})
export class WalletComponent implements OnInit {

  balance: number;

  constructor(
    public modal: NgbActiveModal,
    private toast: ToastrService,
    private http: HttpClient,
    private utils: UtilsService,
  ) { }

  ngOnInit(): void {
    this.init();
  }

  private init() {
    const url = `${this.utils.baseUrl()}/api/wallet`;
     this.http.get(url)
     .subscribe((resp:any)=>{
      this.balance = (JSON.parse(resp.message)).balance;
     },
     err=> {
       console.error(err);
       this.toast.error('Failed to get wallet details.', 'ERROR');
     })
  }

}
