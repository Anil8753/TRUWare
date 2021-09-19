import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import { SpinnerService } from 'src/app/services/spinner.service';
import { UtilsService, WalletEntry } from 'src/app/services/utils.service';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss']
})
export class WalletComponent implements OnInit {

  amount: string;
  refNo: string;

  we: WalletEntry;

  constructor(
    public activeModal: NgbActiveModal,
    private utils: UtilsService,
    private http: HttpClient,
    private toast: ToastrService,
    private spinner: SpinnerService,
  ) { }

  ngOnInit(): void {
    this.init();
  }

  async init() {

    try {
      this.we = await this.utils.getWallet()
    } catch (e){
      console.error(e);
    }
  }

  onLiquidate() {
    this.toast.error('Not implemeted yet.', 'ERROR');
  }

  async onBuyMore() {

    try {

      if (!this.amount || !this.refNo) {
        this.toast.error('Please enter correct amount and ref number.', 'ERROR');
        return;
      }

      this.spinner.show();

      interface PostData {
        amount :string;
        refNo: string;
      }

      const postdata:PostData = {
        amount: this.amount.toString(),
        refNo: this.refNo,
      };

      const url = `${this.utils.baseUrl()}/api/wallet/buy`;
      await this.http.post<any>(url, postdata).toPromise();
      this.amount = '';
      this.refNo = '';
      await this.init();
      this.spinner.hide();
      this.toast.success('Wallet loaded successfully.', 'SUCCESS');

    } catch (e){
      this.spinner.hide();
      this.toast.error('Failed to save the orders data.', 'ERROR');
      console.error(e);
    };
  }
}
