import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FilterbarComponent } from './filterbar.component';

describe('FilterbarComponent', () => {
  let component: FilterbarComponent;
  let fixture: ComponentFixture<FilterbarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FilterbarComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FilterbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
