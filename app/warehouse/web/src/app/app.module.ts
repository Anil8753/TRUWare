import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {HttpClientModule} from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

// Import library module
import { FormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { NgxSpinnerModule } from "ngx-spinner";
import { ToastrModule, ToastContainerModule } from 'ngx-toastr';
import { GoogleMapsModule } from '@angular/google-maps'
import { HomeComponent } from './components/home/home.component';
import { AssetComponent } from './components/asset/asset.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { FormvalidationDirective } from './directives/formvalidation.directive';
import { GmapsComponent } from './components/gmaps/gmaps.component';
import { WarehouseDetailsComponent } from './components/warehouse-details/warehouse-details.component';
import { FooterComponent } from './components/footer/footer.component';
import { RatingsComponent } from './components/ratings/ratings.component';
import { SensorsComponent } from './components/sensors/sensors.component';
import { WalletComponent } from './components/wallet/wallet.component';
import { TransactionsComponent } from './components/transactions/transactions.component';
import { PhotoGalleryComponent } from './components/photo-gallery/photo-gallery.component';
import { VideoGalleryComponent } from './components/video-gallery/video-gallery.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    AssetComponent,
    FormvalidationDirective,
    GmapsComponent,
    WarehouseDetailsComponent,
    FooterComponent,
    RatingsComponent,
    SensorsComponent,
    WalletComponent,
    TransactionsComponent,
    PhotoGalleryComponent,
    VideoGalleryComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    FormsModule,
    NgbModule,
    NgxSpinnerModule,
    ToastrModule,
    GoogleMapsModule,
    FontAwesomeModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot({ positionClass: 'toast-bottom-right', }),
    ToastContainerModule,
  ],
  providers: [],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  bootstrap: [AppComponent]
})
export class AppModule { }
