import { Component, OnInit } from '@angular/core';
import { faTemperatureHigh } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-sensors',
  templateUrl: './sensors.component.html',
  styleUrls: ['./sensors.component.scss']
})
export class SensorsComponent implements OnInit {

  sensorIcon = faTemperatureHigh;

  size = '3x';

  constructor() { }

  ngOnInit(): void {
   ///this.init();
  }

  private init() {
    setInterval(() => {
      this.size = (this.size == '4x') ? '3x' : '4x';
    }, 500);
  }

}
