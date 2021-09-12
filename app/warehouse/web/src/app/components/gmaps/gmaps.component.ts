import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-gmaps',
  templateUrl: './gmaps.component.html',
  styleUrls: ['./gmaps.component.scss']
})
export class GmapsComponent implements OnInit {

  @Input() longitude:number;
  @Input() latitude:number; 

  center: google.maps.LatLngLiteral
  options: google.maps.MapOptions = {
    mapTypeId: 'hybrid',
    zoomControl: true,
    scrollwheel: false,
    disableDoubleClickZoom: true,
    maxZoom: 25,
    minZoom: 8,
  }

  zoom = 15;

  markerOptions: google.maps.MarkerOptions = {draggable: false};
  markerPositions: google.maps.LatLngLiteral[] = [];

  constructor() { }

  ngOnInit(): void {

      this.center = {
        lat: this.latitude,
        lng: this.longitude,
      };

      this.markerPositions.push({ lat: this.latitude, lng: this.longitude});
  }
}
