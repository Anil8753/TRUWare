import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-filterbar',
  templateUrl: './filterbar.component.html',
  styleUrls: ['./filterbar.component.scss']
})
export class FilterbarComponent implements OnInit {

  duration = 90;
  
  constructor() { }

  ngOnInit(): void {
  }

}
