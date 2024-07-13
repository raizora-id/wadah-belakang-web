import Link from 'next/link';
import React from 'react';

const Footer = () => {
    return (
        <>
            <footer className="footer">
                <div className="footer-1">
                    <div className="container">
                        <div className="row">
                            <div className="col-lg-3 width-20">
                                <div className="mb-10"><img src="assets/imgs/template/logo.svg" alt="iori" /></div>
                                <p className="font-md mb-20 color-grey-400">4517 Washington Ave.<br className="d-none d-lg-block" />Manchester, Kentucky 39495</p>
                                <div className="font-md mb-20 color-grey-400"><strong className="font-md-bold">Hours:</strong> 8:00 - 17:00, Mon - Sat</div>
                            </div>
                            <div className="col-lg-3 width-16 mb-30">
                                <h5 className="mb-10 color-brand-1">About Us</h5>
                                <ul className="menu-footer">
                                    <li><Link href="/about">Mission &amp; Vision</Link></li>
                                    <li><Link href="/team">Our Team</Link></li>
                                    <li><Link href="/career">Blog</Link></li>
                                    <li><Link href="#">Partnership</Link></li>
                                    <li><Link href="#">Love Story</Link></li>
                                </ul>
                            </div>
                            <div className="col-lg-3 width-16 mb-30">
                                <h5 className="mb-10 color-brand-1">Products</h5>
                                <ul className="menu-footer">
                                    <li><Link href="#">Aquatic Life</Link></li>
                                    <li><Link href="#">Reptile</Link></li>
                                    <li><Link href="#">Fish Tank</Link></li>
                                    <li><Link href="#">Gundam Plastic</Link></li>
                                    <li><Link href="#">Equipment Aquatic</Link></li>
                                </ul>
                            </div>
                            <div className="col-lg-3 width-16 mb-30">
                                <h5 className="mb-10 color-brand-1">We offer</h5>
                                <ul className="menu-footer">
                                    <li><Link href="#">Aquascape</Link></li>
                                    <li><Link href="#">Custom Fish Tank</Link></li>
                                    <li><Link href="#">Custom Build Gunpla</Link></li>
                                    <li><Link href="#">Custom Diecast</Link></li>
                                </ul>
                            </div>
                            <div className="col-lg-3 width-23">
                                <h5 className="mb-10 color-brand-1">Ecommerce</h5>
                                <div>
                                    <p className="font-sm color-grey-400">Download our Apps and get extra 15% Discount on your first Order…!</p>
                                    <div className="mt-20"><Link className="mr-10" href="#"><img src="assets/imgs/template/appstore.png" alt="iori" /></Link><Link href="#"><img src="assets/imgs/template/google-play.png" alt="iori" /></Link></div>
                                    <p className="font-sm color-grey-400 mt-20 mb-10">Secured Payment Gateways</p>
                                    <div className="mt-15"><Link className="icon-socials icon-facebook" href="#" /><Link className="icon-socials icon-instagram" href="#" /><Link className="icon-socials icon-twitter" href="#" /><Link className="icon-socials icon-linkedin" href="#" /><Link className="icon-socials icon-youtube" href="#" /></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="footer-2">
                    <div className="container">
                        <div className="footer-bottom">
                            <div className="row">
                                <div className="col-lg-6 col-md-12 text-center text-lg-start">
                                    <ul className="menu-bottom">
                                        <li><Link className="font-sm color-grey-300" href="/term-conditions">Privacy policy</Link></li>
                                        <li><Link className="font-sm color-grey-300" href="/term-conditions">Cookies</Link></li>
                                        <li><Link className="font-sm color-grey-300" href="/term-conditions">Terms of service</Link></li>
                                    </ul>
                                </div>
                                <div className="col-lg-6 col-md-12 text-center text-lg-end"><span className="color-grey-300 font-md">©Iori Official 2022. All right reversed.</span></div>
                            </div>
                        </div>
                    </div>
                </div>
            </footer>


        </>
    );
};

export default Footer;