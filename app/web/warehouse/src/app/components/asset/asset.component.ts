import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { SpinnerService } from 'src/app/services/spinner.service';

@Component({
  selector: 'app-asset',
  templateUrl: './asset.component.html',
  styleUrls: ['./asset.component.scss']
})
export class AssetComponent implements OnInit {

  constructor(
    public activeModal: NgbActiveModal,
    private spinner: SpinnerService,
  ) { }

  ngOnInit(): void {
  }

  onSubmit() {
   // alert('Are you sure')
   this.spinner.show();

    setTimeout(()=>{
      this.activeModal.close('submit clicked');
      this.spinner.hide();
    }, 5000);
  }
}
