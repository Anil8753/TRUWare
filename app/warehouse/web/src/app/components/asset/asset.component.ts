import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { faMapPin } from '@fortawesome/free-solid-svg-icons';
import { SpinnerService } from 'src/app/services/spinner.service';
import { ToastrService } from 'ngx-toastr';
import { HttpClient } from '@angular/common/http';
import { UtilsService } from 'src/app/services/utils.service';
import { Asset, GeneralInfo, Postion } from 'src/app/components/asset/asset';

@Component({
  selector: 'app-asset',
  templateUrl: './asset.component.html',
  styleUrls: ['./asset.component.scss']
})
export class AssetComponent implements OnInit {

  @Input() asset: Asset;
  @Input() edit:boolean;

  faMapPin = faMapPin;

  constructor(
    public activeModal: NgbActiveModal,
    private spinner: SpinnerService,
    private toast: ToastrService,
    private http: HttpClient,
    private utils: UtilsService,
  ) { 
    if (!this.edit) {
      this.asset = new Asset();
      this.asset.generalInfo = new GeneralInfo();
      this.asset.postion = new Postion();
    }
  }

  ngOnInit(): void {
  }

  getPosition() {
      navigator.geolocation.getCurrentPosition(resp => {
        this.asset.postion.latitude = resp.coords.latitude;
        this.asset.postion.longitude = resp.coords.longitude;
        },
        err => {
          this.toast.error(`Either location access is blocked or your browser does not support location feature.
          You can enter the location manually.`, 'ERROR');
          console.error(err);
        });
  }

  onDismiss() {
    this.activeModal.close(null);
  }

  onSubmit() {
    if (this.edit)
      this.onUpdate();
    else
      this.onRegister();
  }

  private onRegister() {
    const validation = this.validate();
    if (validation.status === false) {
     this.toast.error(validation.message, 'ERROR');
       return;
    }
 
     this.spinner.show();
     this.asset.id = this.utils.getUUID();
     this.asset.status = parseInt(this.asset.status.toString());
     console.log(this.asset);
 
     const url = `${this.utils.baseUrl()}/api/warehouse`;
 
     this.http.post(url, this.asset)
     .subscribe(res=>{
       this.activeModal.close(this.asset);
       this.spinner.hide();
       this.toast.success('Warehouse data saved succesfully.', 'SUCCESS');
     },
     err=>{
       this.spinner.hide();
       this.toast.error('Failed to save the warehouse data.', 'ERROR');
       console.error(err);
     });
  }

  private onUpdate() {
    const validation = this.validate();
    if (validation.status === false) {
     this.toast.error(validation.message, 'ERROR');
       return;
    }
 
     this.spinner.show();
     this.asset.status = parseInt(this.asset.status.toString());
     console.log(this.asset);
 
     const url = `${this.utils.baseUrl()}/api/warehouse/${this.asset.id}`;
 
     this.http.put(url, this.asset)
     .subscribe(res=>{
       this.activeModal.close(this.asset);
       this.spinner.hide();
       this.toast.success('Warehouse data saved succesfully.', 'SUCCESS');
     },
     err=>{
       this.spinner.hide();
       this.toast.error('Failed to save the warehouse data.', 'ERROR');
       console.error(err);
     });
  }

  private validate() : { status:boolean, message: string}{

    if (this.asset.status == null || this.asset.status == undefined)
      return { status:false, message:`Please enter warehouse status`};

    if (!this.asset.generalInfo.gst)
      return { status:false, message:`Please enter GST number`};

    if (!this.asset.generalInfo.name)
      return { status:false, message:`Please enter warehouse name`};

    if (!this.asset.generalInfo.email)
      return { status:false, message:`Please enter warehouse email`};

    if (!this.asset.generalInfo.phone)
      return { status:false, message:`Please enter warehouse phone number`};

    if (!this.asset.generalInfo.address)
      return { status:false, message:`Please enter warehouse address`};

    if (!this.asset.generalInfo.details)
      return { status:false, message:`Please enter warehouse details`};

    if (!this.asset.generalInfo.totalArea)
      return { status:false, message:`Please enter warehouse total area`};

    if (this.asset.generalInfo.allocatedArea == null)
      return { status:false, message:`Please enter warehouse already occupied space`};

    if (!this.asset.generalInfo.rate)
      return { status:false, message:`Please enter warehouse rate`};

    if (!this.asset.postion.longitude || !this.asset.postion.latitude)
      return { status:false, message:`Please enter warehouse location`};

    return { status: true, message: ''};
  }
}
