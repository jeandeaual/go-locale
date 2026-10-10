//go:build darwin && !ios

#import <Foundation/Foundation.h>
#import <CoreFoundation/CoreFoundation.h>
#include <stdlib.h>
#include <string.h>

bool hasAppTranslations() {
    return [[[NSBundle mainBundle] localizations] count] > 1;
}

const char *preferredLocalization()
{
    if (!hasAppTranslations()) {
        return "";
    }

    NSString *locale = [[[NSBundle mainBundle] preferredLocalizations] firstObject];

    return [locale UTF8String];
}

const char *preferredLocalizations()
{
    if (!hasAppTranslations()) {
        return "";
    }

    NSString *locales = [[[NSBundle mainBundle] preferredLocalizations] componentsJoinedByString:@","];

    return [locales UTF8String];
}

// copyCFString returns a malloc'd UTF-8 copy of s, or NULL. The caller frees it.
static char *copyCFString(CFStringRef s)
{
    if (s == NULL) {
        return NULL;
    }
    CFIndex size = CFStringGetMaximumSizeForEncoding(CFStringGetLength(s), kCFStringEncodingUTF8) + 1;
    char *buf = malloc(size);
    if (buf == NULL) {
        return NULL;
    }
    if (!CFStringGetCString(s, buf, size, kCFStringEncodingUTF8)) {
        free(buf);
        return NULL;
    }
    return buf;
}

// systemLocale returns the identifier of the current user locale (e.g. "en_US"),
// the same value as "defaults read -g AppleLocale". The caller frees it.
char *systemLocale()
{
    CFLocaleRef locale = CFLocaleCopyCurrent();
    if (locale == NULL) {
        return NULL;
    }
    char *out = copyCFString(CFLocaleGetIdentifier(locale));
    CFRelease(locale);
    return out;
}

// systemLanguages returns the user's preferred languages joined by ",",
// the same list as "defaults read -g AppleLanguages". The caller frees it.
char *systemLanguages()
{
    CFArrayRef langs = CFLocaleCopyPreferredLanguages();
    if (langs == NULL) {
        return NULL;
    }
    char *out = NULL;
    if (CFArrayGetCount(langs) > 0) {
        CFStringRef joined = CFStringCreateByCombiningStrings(NULL, langs, CFSTR(","));
        if (joined != NULL) {
            out = copyCFString(joined);
            CFRelease(joined);
        }
    }
    CFRelease(langs);
    return out;
}
