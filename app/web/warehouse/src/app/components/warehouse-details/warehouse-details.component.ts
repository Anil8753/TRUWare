import { Component, Input, OnInit } from '@angular/core';
import { Asset } from '../asset/asset';

@Component({
  selector: 'app-warehouse-details',
  templateUrl: './warehouse-details.component.html',
  styleUrls: ['./warehouse-details.component.scss']
})
export class WarehouseDetailsComponent implements OnInit {

  @Input() asset: Asset;

  constructor() { }

  ngOnInit(): void {
  }

}
