import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent implements OnInit {

  userURL: string;
  shortURL: string;

  constructor(private httpClient: HttpClient) { }

  ngOnInit() {
  }

  shortenURL() {
    this.shortURL = this.userURL + "coco.com"
    console.log(this.shortURL)


    this.httpClient.post("localhost:8080/api/v1/shorten", {"URL": "www.example.com", "Name": ""})

  }

}
