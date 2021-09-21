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
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { FormvalidationDirective } from './directives/formvalidation.directive';
import { GmapsComponent } from './components/gmaps/gmaps.component';
import { FooterComponent } from './components/footer/footer.component';
import { RatingsComponent } from './components/ratings/ratings.component';
import { SensorsComponent } from './components/sensors/sensors.component';
import { FilterbarComponent } from './components/filterbar/filterbar.component';
import { WarehouseListComponent } from './components/warehouse-list/warehouse-list.component';
import { WarehouseCardComponent } from './components/warehouse-card/warehouse-card.component';
import { BookSpaceComponent } from './components/book-space/book-space.component';
import { TransactionsComponent } from './components/transactions/transactions.component';
import { WalletComponent } from './components/wallet/wallet.component';
import { PaymentMethodComponent } from './components/wallet/payment-method/payment-method.component';
import { AccountComponent } from './components/account/account.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    FormvalidationDirective,
    GmapsComponent,
    FooterComponent,
    RatingsComponent,
    SensorsComponent,
    FilterbarComponent,
    WarehouseListComponent,
    WarehouseCardComponent,
    BookSpaceComponent,
    TransactionsComponent,
    WalletComponent,
    PaymentMethodComponent,
    AccountComponent,
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
