import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import Swal from 'ngx-angular8-sweetalert2'
import { SpinnerService } from 'src/app/services/spinner.service';
import { UtilsService } from 'src/app/services/utils.service';

@Component({
  selector: 'app-transactions',
  templateUrl: './transactions.component.html',
  styleUrls: ['./transactions.component.scss']
})
export class TransactionsComponent implements OnInit {

  activeTxns: OrderObject[];
  nonActiveTxns: OrderObject[];

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

  private init() {

    const url = `${this.utils.baseUrl()}/api/orders`;
  
    this.http.get<any>(url)
    .subscribe(res=>{

      const data = JSON.parse(res.message);
      let orders = data as OrderObject[]

      this.activeTxns = orders.filter(o=> o.status == 0)
      this.nonActiveTxns = orders.filter(o=> o.status !== 0)

    },
    err=>{
      this.toast.error('Failed to save the orders data.', 'ERROR');
      console.error(err);
    });
  }

  onCancelBooking(orderId: string) {

    Swal.fire({
      title: 'Are you sure?',
      text: "You won't be able to revert this! This cancellation may cause the associated penalities",
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#3085d6',
      confirmButtonText: 'Yes, cancel the order!',
      customClass: {
        title: 'swl-btn-h',
        content: 'swl-btn-c',
        confirmButton: 'swl-btn-b',
        cancelButton: 'swl-btn-b',
      }
    }).then((result) => {
      if (result.value) {
       this.doCancelBooking(orderId);
      }
    });
  }

  private doCancelBooking(orderId: string) {

    this.spinner.show();

    const data = { comments: 'pleae cancel my booking'};
    const url = `${this.utils.baseUrl()}/api/order/cancel/${orderId}`;
    this.http.put<any>(url, data)
    .subscribe(res=>{
      this.spinner.hide();
      this.toast.success('Succesfully cancelled the order.', 'SUCCESS');
      this.init();
    },
    err=>{
      this.spinner.hide();
      this.toast.error('Failed to cancel the order.', 'ERROR');
      console.error(err);
    });
  }

  getStatusLabel(status: number) {

    if (status == 0)
      return 'Booking is confirmed';

    if (status == 1)
      return 'Booking is active';

    if (status == 2)
      return 'Booking has been rejected';
    
    if (status == 3)
      return 'Booking order has completed';

    if (status == 4)
      return 'You have cancelled the booking';

    if (status == 5)
      return 'You have exceeded the booking tenure';

    return ""
  }
}

class OrderObject {
  id: string;
  type: string;
  warehouseId: string;
  warehouseName: string;
  warehouseGST:string;
  customerId: string;
  date:string;
  value: number;
  rate: number;
  space: number;
  duration: number;
  panalityAfterLimit: number;
  panalityPremature: number;
  status: number;
  conmments: string;
}
