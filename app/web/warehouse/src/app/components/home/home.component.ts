import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { AssetComponent } from '../asset/asset.component';
import { Asset, GeneralInfo, Postion } from 'src/app/components/asset/asset';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  asset: Asset;

  constructor(
    private modalService: NgbModal,
  ) { }

  ngOnInit(): void {
  }

  onClick() {

    const modalRef = this.modalService.open(AssetComponent, { backdrop: 'static'});
    modalRef.componentInstance.name = 'World';

    modalRef.result
    .then((asset: Asset)=>{
      if (asset)
        this.asset = asset;
    })
    .catch(e=>{
      alert(JSON.stringify(e));
    });
  }
}
