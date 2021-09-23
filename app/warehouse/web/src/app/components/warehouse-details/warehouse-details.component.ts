import { Component, Input, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Asset } from '../asset/asset';
import { PhotoGalleryComponent } from '../photo-gallery/photo-gallery.component';
import { VideoGalleryComponent } from '../video-gallery/video-gallery.component';

@Component({
  selector: 'app-warehouse-details',
  templateUrl: './warehouse-details.component.html',
  styleUrls: ['./warehouse-details.component.scss']
})
export class WarehouseDetailsComponent implements OnInit {

  @Input() asset: Asset;

  constructor(
    private modalService: NgbModal,
  ) { }

  ngOnInit(): void {
  }

  onPhotoGallary(whID: string) {
    this.modalService.open(PhotoGalleryComponent, { backdrop: 'static', size: 'lg'});
  }

  onVideoGallary(whID: string) {
    this.modalService.open(VideoGalleryComponent, { backdrop: 'static', size: 'lg'});
  }
}
