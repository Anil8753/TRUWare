import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { AssetComponent } from '../asset/asset.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  constructor(
    private modalService: NgbModal,
  ) { }

  ngOnInit(): void {
  }

  onClick() {

    const modalRef = this.modalService.open(AssetComponent);
    modalRef.componentInstance.name = 'World';

    modalRef.result
    .then(reson=>{
      alert(reson);
    })
    .catch(e=>{
      alert(JSON.stringify(e));
    });
  }
}
