import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent implements OnInit {

  userURL: string;
  shortURL: string;

  constructor() { }

  ngOnInit() {
  }

  shortenURL() {
    this.shortURL = this.userURL + "coco.com"
    console.log(this.shortURL)
  }

}
