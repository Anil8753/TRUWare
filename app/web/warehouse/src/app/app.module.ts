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
import { HomeComponent } from './components/home/home.component';
import { AssetComponent } from './components/asset/asset.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { FormvalidationDirective } from './directives/formvalidation.directive';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    AssetComponent,
    FormvalidationDirective,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    FormsModule,
    NgbModule,
    NgxSpinnerModule,
    ToastrModule,
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
