import { Component, OnInit } from '@angular/core';
import { faFilter} from '@fortawesome/free-solid-svg-icons';
import { MessageService, MessagePayload, MessageType } from 'src/app/services/message.service';

@Component({
  selector: 'app-filterbar',
  templateUrl: './filterbar.component.html',
  styleUrls: ['./filterbar.component.scss']
})
export class FilterbarComponent implements OnInit {

  duration = 90;
  iconClearFilter = faFilter;
  hasFilter = false;
  currFilter = '';
  
  constructor(
    private msgService: MessageService,
  ) { }

  ngOnInit(): void {
  }

  onClearFilter() {
    const payload = new MessagePayload();
    payload.type = MessageType.ClearFilter;
    this.msgService.sendMessage(payload);
    this.hasFilter = false;
    this.fromRate = 0;
    this.toRate = 0;
    this.space = 0;
    this.address = '';
    this.currFilter = '';
  }

  onFilterByRating(n:number) {
    const payload = new MessagePayload();
    payload.type = MessageType.RatingsFilter;
    payload.data = n;
    this.msgService.sendMessage(payload);
    this.hasFilter = true;
  }

  toRate:number;
  fromRate:number;

  onFilterByRateRange() {
    const payload = new MessagePayload();
    payload.type = MessageType.RateFilter;
    payload.data = { to:this.toRate, from:this.fromRate };
    this.msgService.sendMessage(payload);
    this.hasFilter = true;
    this.currFilter = 'Filtered by rate';
  }

  space:number;

  onFilterBySpace() {
    const payload = new MessagePayload();
    payload.type = MessageType.SpaceFilter;
    payload.data = this.space;
    this.msgService.sendMessage(payload);
    this.hasFilter = true;
    this.currFilter = 'Filtered by space';
  }

  address:string;

  onFilterByAddress() {

    if (!this.address || this.address === '')
      return;

    const payload = new MessagePayload();
    payload.type = MessageType.LocationFilter;
    payload.data = this.address;
    this.msgService.sendMessage(payload);
    this.hasFilter = true;
    this.currFilter = 'Filtered by location';
  }
}
