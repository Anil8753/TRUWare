import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import { UtilsService } from 'src/app/services/utils.service';

@Component({
  selector: 'app-transactions',
  templateUrl: './transactions.component.html',
  styleUrls: ['./transactions.component.scss']
})
export class TransactionsComponent implements OnInit {

  orders : any[];

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
    const url = `${this.utils.baseUrl()}/api/orders`;
     this.http.get(url)
     .subscribe((resp:any)=>{
      this.orders = (JSON.parse(resp.message));
     },
     err=> {
       console.error(err);
       this.toast.error('Failed to get orders details.', 'ERROR');
     })
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
      return 'Booking is cancelled by customer';

    if (status == 5)
      return 'You have exceeded the booking tenure';

    return ""
  }
}
