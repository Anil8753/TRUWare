import { HttpClient } from '@angular/common/http';
import { Component, OnInit, Input } from '@angular/core';
import { NgbActiveModal, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import { SpinnerService } from 'src/app/services/spinner.service';
import { UtilsService, WalletEntry } from 'src/app/services/utils.service';
import { Order } from '../warehouse-card/order';
import { Warehouse } from '../warehouse-card/warehouse';

@Component({
  selector: 'app-book-space',
  templateUrl: './book-space.component.html',
  styleUrls: ['./book-space.component.scss']
})
export class BookSpaceComponent implements OnInit {

  @Input() warehouse: Warehouse;

  order: Order;
  duration: number;
  comments: string;
  space: number;
  cost: number;

  constructor(
    public activeModal: NgbActiveModal,
    private utils: UtilsService,
    private http: HttpClient,
    private spinner: SpinnerService,
    private toast: ToastrService,
  ) { }

  ngOnInit(): void {
    this.duration = 0;
    this.space = 0;
    this.cost = 0;

    if (!this.warehouse.generalInfo.panalityPremature) {
      this.warehouse.generalInfo.panalityPremature = 0;
    }

    if (!this.warehouse.generalInfo.panalityAfterLimit) {
      this.warehouse.generalInfo.panalityAfterLimit = 0;
    }
  }

  async bookSpace() {
    this.order = new Order;
    this.order.id = this.utils.getUUID();
    this.order.comments = this.comments;
    this.order.duration = this.duration;
    this.order.space = this.space;
    this.order.value = this.space * this.duration *this.warehouse.generalInfo.rate;
    this.order.warehouseId = this.warehouse.id;
    console.log(this.order);

    const validation = this.validate();
    if (validation.status === false) {
      this.toast.error(validation.message, 'ERROR');
        return;
     }
  
      this.spinner.show();
     
      const we = await this.utils.getWallet();
      if (!we) {
        this.spinner.hide();
        this.toast.error('Wallet is not available.', 'ERROR');
        return;
      }

      if (we.balance < this.order.value) {
        this.spinner.hide();
        this.toast.error('Not enough wallet balance, please reload your wallet.', 'ERROR');
        return;
      }

      const url = `${this.utils.baseUrl()}/api/order`;
  
      this.http.post(url, this.order)
      .subscribe(res=>{
        this.activeModal.close(this.order);
        this.spinner.hide();
        this.toast.success('Warehouse data saved succesfully.', 'SUCCESS');

        setTimeout(()=>{ window.location.reload(); }, 500);
      },
      err=>{
        this.spinner.hide();
        this.toast.error('Failed to save the warehouse data.', 'ERROR');
        console.error(err);
      });
  }

  private validate() : { status:boolean, message: string}{

    if (this.order.space <= 0)
      return { status:false, message:`Please enter the valid warehouse space`};

    if (!this.order.duration)
      return { status:false, message:`Please enter valid duration`};

    if (this.order.duration > 365)
      return { status:false, message:`Cannot book the space for more than 365 days`};

    const available = this.warehouse.generalInfo.totalArea - this.warehouse.generalInfo.allocatedArea;

    if (this.order.space > available)
      return { status:false, message:`Cannot book the space more than available space`};

    
    return { status: true, message: ''};
  }
}
