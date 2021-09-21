import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ToastrService } from 'ngx-toastr';
import { SpinnerService } from 'src/app/services/spinner.service';
import { UtilsService, RegistrationData } from 'src/app/services/utils.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.scss']
})
export class AccountComponent implements OnInit {

  registered = false;
  regData: RegistrationData;

  constructor(
    public activeModal: NgbActiveModal,
    private utils: UtilsService,
    private http: HttpClient,
    private spinner: SpinnerService,
    private toast: ToastrService,
  ) { 
    this.regData = new RegistrationData;
    this.regData.name = '';
    this.regData.contact = '';
    this.regData.email = '';
    this.regData.address = '';
    this.regData.gst = '';
  }

  ngOnInit(): void {
    this.init();
  }

  private async init() {

    const resp = await this.utils.getRegistration();
    if (resp == null) {
      this.registered = false;
    } else {
      this.regData = resp;
      this.registered = true;
    }
  }

  register() {

    const v = this.validate();
    if (!v.status) {
      this.toast.error(v.msg, 'ERROR');
      return;
    }

    this.spinner.show();
    this.regData.id = this.utils.getUUID();
    const url = `${this.utils.baseUrl()}/api/registeration`;
  
    this.http.post(url, this.regData)
    .subscribe(res=>{
      this.activeModal.close('');
      this.spinner.hide();
      this.toast.success('Registered succesfully.', 'SUCCESS');
    },
    err=>{
      this.spinner.hide();
      this.toast.error('Failed to register.', 'ERROR');
      console.error(err);
    });
  }

  update() {

    const v = this.validate();
    if (!v.status) {
      this.toast.error(v.msg, 'ERROR');
      return;
    }

    this.spinner.show();
    const url = `${this.utils.baseUrl()}/api/registeration`;
  
    this.http.put(url, this.regData)
    .subscribe(res=>{
      this.activeModal.close('');
      this.spinner.hide();
      this.toast.success('Updated succesfully.', 'SUCCESS');
    },
    err=>{
      this.spinner.hide();
      this.toast.error('Failed to update.', 'ERROR');
      console.error(err);
    });

  }

  private validate() {

    if (!this.regData.name || this.regData.name === '') {
      return { status: false, msg: 'Please enter name'};
    }

    if (!this.regData.address || this.regData.address === '') {
      return { status: false, msg: 'Please enter address'};
    }

    if (!this.regData.contact || this.regData.contact === '') {
      return { status: false, msg: 'Please enter the contact number'};
    }

    if (!this.regData.email || this.regData.email === '') {
      return { status: false, msg: 'Please enter the email address'};
    }

    if (!this.regData.gst || this.regData.gst === '') {
      return { status: false, msg: 'Please enter the GST number'};
    }

    return { status: true, msg: ''};
  }
}

