import { TestBed } from '@angular/core/testing';

import { GooglemapService } from './googlemap.service';

describe('GooglemapService', () => {
  let service: GooglemapService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GooglemapService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
