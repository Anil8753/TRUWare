import { Component, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { UtilsService } from 'src/app/services/utils.service';

@Component({
  selector: 'app-photo-gallery',
  templateUrl: './photo-gallery.component.html',
  styleUrls: ['./photo-gallery.component.scss']
})
export class PhotoGalleryComponent implements OnInit {

  images = [];

  constructor(
    public modal: NgbActiveModal,
  ) { 
  }

  ngOnInit(): void {
    this.images.push('assets/whgalley/1.png')
    this.images.push('assets/whgalley/2.png')
    this.images.push('assets/whgalley/3.png')
    this.images.push('assets/whgalley/4.png')
  }

}
