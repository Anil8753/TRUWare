import { Component, Input, OnInit } from '@angular/core';
import { faStar} from '@fortawesome/free-solid-svg-icons';
import { faMapMarkerAlt, faCloudRain, faTemperatureLow} from '@fortawesome/free-solid-svg-icons';
import { NgbActiveModal, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { BookSpaceComponent } from '../book-space/book-space.component';
import { Warehouse } from './warehouse';

@Component({
  selector: 'app-warehouse-card',
  templateUrl: './warehouse-card.component.html',
  styleUrls: ['./warehouse-card.component.scss']
})
export class WarehouseCardComponent implements OnInit {

  @Input() warehouse: Warehouse;

  starIcon = faStar;

  iconHumidity = faCloudRain;
  iconTemp = faTemperatureLow;
  iconLocation = faMapMarkerAlt;

  constructor(
    private modalService: NgbModal
  ) { }

  ngOnInit(): void {
  }

  counter(i: number) {
    return new Array(i);
  }

  openMap(content) {
    this.modalService.open(content);
  }

  bookSpace() {
    const modalRef = this.modalService.open(BookSpaceComponent, { backdrop: 'static', size: 'lg'});
    modalRef.componentInstance.warehouse = this.warehouse;
  }
}
