import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BookSpaceComponent } from './book-space.component';

describe('BookSpaceComponent', () => {
  let component: BookSpaceComponent;
  let fixture: ComponentFixture<BookSpaceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BookSpaceComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BookSpaceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
